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

    public struct Point9100
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid9100
    {
        public static Point9100 Make(int left, int right)
        {
            return new Point9100 { X = left, Y = right };
        }
    }
    public interface IHandler9101
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9101(string Label, int Payload);

    public sealed class Handler9101 : IHandler9101
    {
        public void Handle(int value) { }
    }
    public static int Scale9102<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public sealed class SlotMap9103
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
