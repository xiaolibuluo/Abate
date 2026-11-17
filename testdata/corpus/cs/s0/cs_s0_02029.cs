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

    public static int Guard18930(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 73;
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
    public struct Point18931
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid18931
    {
        public static Point18931 Make(int left, int right)
        {
            return new Point18931 { X = left, Y = right };
        }
    }
    public struct Point18932
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 87;
    }

    public static class Grid18932
    {
        public static Point18932 Make(int left, int right)
        {
            return new Point18932 { X = left, Y = right };
        }
    }
    public delegate void Notifier18933(int value);

    public sealed class Source18933
    {
        public event Notifier18933 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point18934
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid18934
    {
        public static Point18934 Make(int left, int right)
        {
            return new Point18934 { X = left, Y = right };
        }
    }
    public static int Guard18935(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 11;
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
}
