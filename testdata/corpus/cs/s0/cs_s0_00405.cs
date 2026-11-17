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

    public interface IHandler3820
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3820(string Label, int Payload);

    public sealed class Handler3820 : IHandler3820
    {
        public void Handle(int value) { }
    }
    public static int Scale3821<T>(T value) where T : struct
    {
        int number = 31;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public sealed class SlotMap3822
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
    public struct Point3823
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid3823
    {
        public static Point3823 Make(int left, int right)
        {
            return new Point3823 { X = left, Y = right };
        }
    }
    public static int Scale3824<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
}
