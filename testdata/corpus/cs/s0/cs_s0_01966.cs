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

    public delegate void Notifier18320(int value);

    public sealed class Source18320
    {
        public event Notifier18320 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18321<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public struct Point18322
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 70;
    }

    public static class Grid18322
    {
        public static Point18322 Make(int left, int right)
        {
            return new Point18322 { X = left, Y = right };
        }
    }
    public static int Guard18323(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 77;
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
    public static int Scale18324<T>(T value) where T : struct
    {
        int number = 84;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 67; }
        return number;
    }
}
