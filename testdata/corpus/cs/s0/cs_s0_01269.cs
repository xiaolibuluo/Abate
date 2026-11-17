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

    public delegate void Notifier11890(int value);

    public sealed class Source11890
    {
        public event Notifier11890 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard11891(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 46;
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
    public struct Point11892
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid11892
    {
        public static Point11892 Make(int left, int right)
        {
            return new Point11892 { X = left, Y = right };
        }
    }
}
