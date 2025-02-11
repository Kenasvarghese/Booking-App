import Footer from "./components/Footer";
import Header from "./components/Header";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  console.log("test");

  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      {/* <div className="container mx-auto px-4  space-y-4"> */}

      <main className="flex-grow container mx-auto">{children}</main>
      <Footer />
    </div>
  );
}
