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

    public sealed class SlotMap24630
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
    public interface IHandler24631
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24631(string Label, int Payload);

    public sealed class Handler24631 : IHandler24631
    {
        public void Handle(int value) { }
    }
    public static int Scale24632<T>(T value) where T : struct
    {
        int number = 28;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public struct Point24633
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid24633
    {
        public static Point24633 Make(int left, int right)
        {
            return new Point24633 { X = left, Y = right };
        }
    }
}
