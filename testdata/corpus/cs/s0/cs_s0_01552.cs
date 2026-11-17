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

    public sealed class SlotMap14520
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
    public static int Guard14521(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 70;
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
    public delegate void Notifier14522(int value);

    public sealed class Source14522
    {
        public event Notifier14522 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14523
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14523(string Label, int Payload);

    public sealed class Handler14523 : IHandler14523
    {
        public void Handle(int value) { }
    }
    public interface IHandler14524
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14524(string Label, int Payload);

    public sealed class Handler14524 : IHandler14524
    {
        public void Handle(int value) { }
    }
}
