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

    public interface IHandler750
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope750(string Label, int Payload);

    public sealed class Handler750 : IHandler750
    {
        public void Handle(int value) { }
    }
    public struct Point751
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 17;
    }

    public static class Grid751
    {
        public static Point751 Make(int left, int right)
        {
            return new Point751 { X = left, Y = right };
        }
    }
    public interface IHandler752
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope752(string Label, int Payload);

    public sealed class Handler752 : IHandler752
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap753
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
