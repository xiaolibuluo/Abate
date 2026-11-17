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

    public static int Guard16900(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 53;
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
    public delegate void Notifier16901(int value);

    public sealed class Source16901
    {
        public event Notifier16901 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16902
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 67;
    }

    public static class Grid16902
    {
        public static Point16902 Make(int left, int right)
        {
            return new Point16902 { X = left, Y = right };
        }
    }
    public struct Point16903
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid16903
    {
        public static Point16903 Make(int left, int right)
        {
            return new Point16903 { X = left, Y = right };
        }
    }
    public static int Guard16904(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 81;
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
