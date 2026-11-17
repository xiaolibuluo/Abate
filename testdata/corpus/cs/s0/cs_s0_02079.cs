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

    public interface IHandler19390
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19390(string Label, int Payload);

    public sealed class Handler19390 : IHandler19390
    {
        public void Handle(int value) { }
    }
    public struct Point19391
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 96;
    }

    public static class Grid19391
    {
        public static Point19391 Make(int left, int right)
        {
            return new Point19391 { X = left, Y = right };
        }
    }
    public delegate void Notifier19392(int value);

    public sealed class Source19392
    {
        public event Notifier19392 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19393(int value);

    public sealed class Source19393
    {
        public event Notifier19393 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point19394
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid19394
    {
        public static Point19394 Make(int left, int right)
        {
            return new Point19394 { X = left, Y = right };
        }
    }
    public sealed class SlotMap19395
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
