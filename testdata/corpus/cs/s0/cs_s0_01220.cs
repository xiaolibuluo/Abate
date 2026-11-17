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

    public static int Guard11450(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 49;
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
    public struct Point11451
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 56;
    }

    public static class Grid11451
    {
        public static Point11451 Make(int left, int right)
        {
            return new Point11451 { X = left, Y = right };
        }
    }
    public delegate void Notifier11452(int value);

    public sealed class Source11452
    {
        public event Notifier11452 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point11453
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 70;
    }

    public static class Grid11453
    {
        public static Point11453 Make(int left, int right)
        {
            return new Point11453 { X = left, Y = right };
        }
    }
}
