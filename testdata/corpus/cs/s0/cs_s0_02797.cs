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

    public sealed class SlotMap26160
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
    public static int Scale26161<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public static int Scale26162<T>(T value) where T : struct
    {
        int number = 77;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public struct Point26163
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 84;
    }

    public static class Grid26163
    {
        public static Point26163 Make(int left, int right)
        {
            return new Point26163 { X = left, Y = right };
        }
    }
    public interface IHandler26164
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26164(string Label, int Payload);

    public sealed class Handler26164 : IHandler26164
    {
        public void Handle(int value) { }
    }
}
