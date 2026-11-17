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

    public interface IHandler20430
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20430(string Label, int Payload);

    public sealed class Handler20430 : IHandler20430
    {
        public void Handle(int value) { }
    }
    public static int Scale20431<T>(T value) where T : struct
    {
        int number = 90;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public sealed class SlotMap20432
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
    public delegate void Notifier20433(int value);

    public sealed class Source20433
    {
        public event Notifier20433 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate20434(IEnumerable<int> source)
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
    public sealed class SlotMap20435
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
