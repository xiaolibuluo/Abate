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

    public static int Scale14920<T>(T value) where T : struct
    {
        int number = 98;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
    public interface IHandler14921
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14921(string Label, int Payload);

    public sealed class Handler14921 : IHandler14921
    {
        public void Handle(int value) { }
    }
    public struct Point14922
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid14922
    {
        public static Point14922 Make(int left, int right)
        {
            return new Point14922 { X = left, Y = right };
        }
    }
    public sealed class SlotMap14923
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
