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

    public struct Point16970
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 47;
    }

    public static class Grid16970
    {
        public static Point16970 Make(int left, int right)
        {
            return new Point16970 { X = left, Y = right };
        }
    }
    public sealed class SlotMap16971
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
    public interface IHandler16972
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16972(string Label, int Payload);

    public sealed class Handler16972 : IHandler16972
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap16973
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
