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

    public static int Scale9690<T>(T value) where T : struct
    {
        int number = 89;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public sealed class SlotMap9691
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
    public interface IHandler9692
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9692(string Label, int Payload);

    public sealed class Handler9692 : IHandler9692
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier9693(int value);

    public sealed class Source9693
    {
        public event Notifier9693 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale9694<T>(T value) where T : struct
    {
        int number = 20;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 65; }
        return number;
    }
    public struct Point9695
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid9695
    {
        public static Point9695 Make(int left, int right)
        {
            return new Point9695 { X = left, Y = right };
        }
    }
}
