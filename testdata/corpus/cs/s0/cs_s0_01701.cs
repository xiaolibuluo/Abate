using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public interface IHandler15890
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15890(string Label, int Payload);

    public sealed class Handler15890 : IHandler15890
    {
        public void Handle(int value) { }
    }
    public static int Guard15891(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 8;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public interface IHandler15892
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15892(string Label, int Payload);

    public sealed class Handler15892 : IHandler15892
    {
        public void Handle(int value) { }
    }
    public struct Point15893
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 22;
    }

    public static class Grid15893
    {
        public static Point15893 Make(int left, int right)
        {
            return new Point15893 { X = left, Y = right };
        }
    }
}
