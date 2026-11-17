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

    public struct Point17340
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid17340
    {
        public static Point17340 Make(int left, int right)
        {
            return new Point17340 { X = left, Y = right };
        }
    }
    public interface IHandler17341
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17341(string Label, int Payload);

    public sealed class Handler17341 : IHandler17341
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap17342
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
    public interface IHandler17343
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17343(string Label, int Payload);

    public sealed class Handler17343 : IHandler17343
    {
        public void Handle(int value) { }
    }
}
