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

    public struct Point19310
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 82;
    }

    public static class Grid19310
    {
        public static Point19310 Make(int left, int right)
        {
            return new Point19310 { X = left, Y = right };
        }
    }
    public delegate void Notifier19311(int value);

    public sealed class Source19311
    {
        public event Notifier19311 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard19312(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 96;
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
    public struct Point19313
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid19313
    {
        public static Point19313 Make(int left, int right)
        {
            return new Point19313 { X = left, Y = right };
        }
    }
}
