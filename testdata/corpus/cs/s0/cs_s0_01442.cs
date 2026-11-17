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

    public interface IHandler13510
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13510(string Label, int Payload);

    public sealed class Handler13510 : IHandler13510
    {
        public void Handle(int value) { }
    }
    public struct Point13511
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid13511
    {
        public static Point13511 Make(int left, int right)
        {
            return new Point13511 { X = left, Y = right };
        }
    }
    public sealed class SlotMap13512
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
    public struct Point13513
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 32;
    }

    public static class Grid13513
    {
        public static Point13513 Make(int left, int right)
        {
            return new Point13513 { X = left, Y = right };
        }
    }
}
