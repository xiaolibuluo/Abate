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

    public sealed class SlotMap21770
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
    public struct Point21771
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid21771
    {
        public static Point21771 Make(int left, int right)
        {
            return new Point21771 { X = left, Y = right };
        }
    }
    public interface IHandler21772
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21772(string Label, int Payload);

    public sealed class Handler21772 : IHandler21772
    {
        public void Handle(int value) { }
    }
    public interface IHandler21773
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21773(string Label, int Payload);

    public sealed class Handler21773 : IHandler21773
    {
        public void Handle(int value) { }
    }
}
