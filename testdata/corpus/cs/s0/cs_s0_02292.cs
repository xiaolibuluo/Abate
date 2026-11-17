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

    public interface IHandler21390
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21390(string Label, int Payload);

    public sealed class Handler21390 : IHandler21390
    {
        public void Handle(int value) { }
    }
    public static int Scale21391<T>(T value) where T : struct
    {
        int number = 77;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 48; }
        return number;
    }
    public static int Guard21392(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 84;
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
    public delegate void Notifier21393(int value);

    public sealed class Source21393
    {
        public event Notifier21393 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21394<T>(T value) where T : struct
    {
        int number = 98;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public sealed class SlotMap21395
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
