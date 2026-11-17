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

    public delegate void Notifier20130(int value);

    public sealed class Source20130
    {
        public event Notifier20130 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard20131(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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
    public struct Point20132
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 95;
    }

    public static class Grid20132
    {
        public static Point20132 Make(int left, int right)
        {
            return new Point20132 { X = left, Y = right };
        }
    }
}
