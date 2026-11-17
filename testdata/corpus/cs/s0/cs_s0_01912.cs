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

    public sealed class SlotMap17850
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
    public interface IHandler17851
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17851(string Label, int Payload);

    public sealed class Handler17851 : IHandler17851
    {
        public void Handle(int value) { }
    }
    public static int Scale17852<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public struct Point17853
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid17853
    {
        public static Point17853 Make(int left, int right)
        {
            return new Point17853 { X = left, Y = right };
        }
    }
}
