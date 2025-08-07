import Link from "next/link";
import { redirect } from "next/navigation";

import { auth } from "~/server/auth";

export default async function Home() {
  const session = await auth();

  // Redirect authenticated users to dashboard
  if (session?.user) {
    redirect("/dashboard");
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-gradient-to-b from-blue-900 to-blue-950 text-white">
      <div className="container flex flex-col items-center justify-center gap-12 px-4 py-16">
        <div className="text-center">
          <h1 className="text-6xl font-extrabold tracking-tight sm:text-7xl">
            <span className="text-blue-400">Elastic</span>Page
          </h1>
          <p className="mt-6 text-xl text-blue-200">
            Transform static textbooks into interactive learning experiences with AI-powered summaries and quizzes
          </p>
        </div>

        <div className="grid grid-cols-1 gap-6 sm:grid-cols-3 md:gap-8">
          <div className="flex flex-col gap-4 rounded-xl bg-white/10 p-6 backdrop-blur-sm">
            <h3 className="text-2xl font-bold text-blue-300">📚 Upload PDFs</h3>
            <p className="text-blue-100">
              Upload your textbook PDFs and let our AI extract and organize the content for interactive learning.
            </p>
          </div>
          
          <div className="flex flex-col gap-4 rounded-xl bg-white/10 p-6 backdrop-blur-sm">
            <h3 className="text-2xl font-bold text-blue-300">🤖 AI Summaries</h3>
            <p className="text-blue-100">
              Get intelligent summaries and key insights generated from your textbook content using advanced AI.
            </p>
          </div>
          
          <div className="flex flex-col gap-4 rounded-xl bg-white/10 p-6 backdrop-blur-sm">
            <h3 className="text-2xl font-bold text-blue-300">📊 Track Progress</h3>
            <p className="text-blue-100">
              Monitor your learning progress with interactive quizzes and detailed analytics.
            </p>
          </div>
        </div>

        <div className="flex flex-col items-center gap-4">
          <Link
            href="/api/auth/signin"
            className="rounded-full bg-blue-600 px-8 py-4 text-lg font-semibold text-white transition hover:bg-blue-700"
          >
            Get Started
          </Link>
          <p className="text-sm text-blue-300">
            Sign in to start transforming your textbooks
          </p>
        </div>
      </div>
    </main>
  );
}
