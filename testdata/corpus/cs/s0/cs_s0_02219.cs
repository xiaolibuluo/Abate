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

    public interface IHandler20720
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20720(string Label, int Payload);

    public sealed class Handler20720 : IHandler20720
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap20721
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
    public struct Point20722
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid20722
    {
        public static Point20722 Make(int left, int right)
        {
            return new Point20722 { X = left, Y = right };
        }
    }
}
