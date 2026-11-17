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

    public static int Guard15790(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 65;
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
    public struct Point15791
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 72;
    }

    public static class Grid15791
    {
        public static Point15791 Make(int left, int right)
        {
            return new Point15791 { X = left, Y = right };
        }
    }
    public interface IHandler15792
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15792(string Label, int Payload);

    public sealed class Handler15792 : IHandler15792
    {
        public void Handle(int value) { }
    }
}
