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

    public sealed class SlotMap25670
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
    public interface IHandler25671
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25671(string Label, int Payload);

    public sealed class Handler25671 : IHandler25671
    {
        public void Handle(int value) { }
    }
    public struct Point25672
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 22;
    }

    public static class Grid25672
    {
        public static Point25672 Make(int left, int right)
        {
            return new Point25672 { X = left, Y = right };
        }
    }
    public static int Scale25673<T>(T value) where T : struct
    {
        int number = 29;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
}
