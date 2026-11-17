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

    public sealed class SlotMap19450
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
    public interface IHandler19451
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19451(string Label, int Payload);

    public sealed class Handler19451 : IHandler19451
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier19452(int value);

    public sealed class Source19452
    {
        public event Notifier19452 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard19453(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 91;
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
    public interface IHandler19454
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19454(string Label, int Payload);

    public sealed class Handler19454 : IHandler19454
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate19455(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
}
