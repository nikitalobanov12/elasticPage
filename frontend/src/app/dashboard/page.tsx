"use client";

import Link from "next/link";
import { useSession } from "next-auth/react";
import { redirect } from "next/navigation";

import { useTextbooks } from "~/hooks/use-textbooks";

export default function Dashboard() {
  const { data: session, status } = useSession();

  if (status === "loading") {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="text-lg">Loading...</div>
      </div>
    );
  }

  if (!session?.user) {
    redirect("/");
  }

  const { data: textbooks, isLoading, error } = useTextbooks(session.user.id);

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow">
        <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between">
            <h1 className="text-3xl font-bold tracking-tight text-gray-900">
              Dashboard
            </h1>
            <div className="flex items-center gap-4">
              <span className="text-sm text-gray-600">
                Welcome, {session.user.name}
              </span>
              <Link
                href="/api/auth/signout"
                className="rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white hover:bg-red-700"
              >
                Sign Out
              </Link>
            </div>
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
        <div className="mb-8 flex items-center justify-between">
          <h2 className="text-2xl font-bold text-gray-900">Your Textbooks</h2>
          <Link
            href="/upload"
            className="rounded-md bg-blue-600 px-4 py-2 text-sm font-semibold text-white hover:bg-blue-700"
          >
            Upload New Textbook
          </Link>
        </div>

        {isLoading && (
          <div className="flex items-center justify-center py-12">
            <div className="text-lg text-gray-600">Loading textbooks...</div>
          </div>
        )}

        {error && (
          <div className="rounded-md bg-red-50 p-4">
            <div className="text-sm text-red-800">
              Error loading textbooks: {error.message}
            </div>
          </div>
        )}

        {textbooks && textbooks.length === 0 && (
          <div className="text-center py-12">
            <div className="text-gray-500">
              <svg
                className="mx-auto h-12 w-12 text-gray-400"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.746 0 3.332.477 4.5 1.253v13C19.832 18.477 18.246 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
                />
              </svg>
              <h3 className="mt-2 text-sm font-medium text-gray-900">No textbooks</h3>
              <p className="mt-1 text-sm text-gray-500">
                Get started by uploading your first textbook.
              </p>
              <div className="mt-6">
                <Link
                  href="/upload"
                  className="inline-flex items-center rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white hover:bg-blue-700"
                >
                  Upload Textbook
                </Link>
              </div>
            </div>
          </div>
        )}

        {textbooks && textbooks.length > 0 && (
          <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {textbooks.map((textbook) => (
              <div
                key={textbook.id}
                className="overflow-hidden rounded-lg bg-white shadow hover:shadow-md transition-shadow"
              >
                <div className="p-6">
                  <h3 className="text-lg font-medium text-gray-900 mb-2">
                    {textbook.title}
                  </h3>
                  {textbook.description && (
                    <p className="text-sm text-gray-600 mb-4">
                      {textbook.description}
                    </p>
                  )}
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-gray-500">
                      {new Date(textbook.created_at).toLocaleDateString()}
                    </span>
                    <Link
                      href={`/textbooks/${textbook.id}`}
                      className="text-sm font-medium text-blue-600 hover:text-blue-700"
                    >
                      View →
                    </Link>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </main>
    </div>
  );
}