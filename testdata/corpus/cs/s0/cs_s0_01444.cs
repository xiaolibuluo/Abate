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

    public struct Point13530
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid13530
    {
        public static Point13530 Make(int left, int right)
        {
            return new Point13530 { X = left, Y = right };
        }
    }
    public sealed class SlotMap13531
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
    public static int Guard13532(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 51;
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
    public struct Point13533
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 58;
    }

    public static class Grid13533
    {
        public static Point13533 Make(int left, int right)
        {
            return new Point13533 { X = left, Y = right };
        }
    }
}
