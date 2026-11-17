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

    public interface IHandler17360
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17360(string Label, int Payload);

    public sealed class Handler17360 : IHandler17360
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap17361
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
    public struct Point17362
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid17362
    {
        public static Point17362 Make(int left, int right)
        {
            return new Point17362 { X = left, Y = right };
        }
    }
    public sealed class SlotMap17363
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
