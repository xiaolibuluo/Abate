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

    public sealed class SlotMap440
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
    public async Task<int> Aggregate441(IEnumerable<int> source)
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
    public interface IHandler442
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope442(string Label, int Payload);

    public sealed class Handler442 : IHandler442
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier443(int value);

    public sealed class Source443
    {
        public event Notifier443 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler444
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope444(string Label, int Payload);

    public sealed class Handler444 : IHandler444
    {
        public void Handle(int value) { }
    }
}
