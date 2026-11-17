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

    public sealed class SlotMap7710
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
    public delegate void Notifier7711(int value);

    public sealed class Source7711
    {
        public event Notifier7711 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler7712
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7712(string Label, int Payload);

    public sealed class Handler7712 : IHandler7712
    {
        public void Handle(int value) { }
    }
    public static int Guard7713(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 58;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public struct Point7714
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 65;
    }

    public static class Grid7714
    {
        public static Point7714 Make(int left, int right)
        {
            return new Point7714 { X = left, Y = right };
        }
    }
}
