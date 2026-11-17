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

    public static int Guard2890(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 76;
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
    public static int Scale2891<T>(T value) where T : struct
    {
        int number = 83;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public struct Point2892
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 90;
    }

    public static class Grid2892
    {
        public static Point2892 Make(int left, int right)
        {
            return new Point2892 { X = left, Y = right };
        }
    }
    public delegate void Notifier2893(int value);

    public sealed class Source2893
    {
        public event Notifier2893 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap2894
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
    public static int Scale2895<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
}
