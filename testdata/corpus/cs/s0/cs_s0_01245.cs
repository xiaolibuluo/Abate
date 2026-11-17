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

    public static int Guard11670(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 44;
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
    public static int Scale11671<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public static int Guard11672(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 58;
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
    public delegate void Notifier11673(int value);

    public sealed class Source11673
    {
        public event Notifier11673 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point11674
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 72;
    }

    public static class Grid11674
    {
        public static Point11674 Make(int left, int right)
        {
            return new Point11674 { X = left, Y = right };
        }
    }
    public sealed class SlotMap11675
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
}
