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

    public async Task<int> Aggregate24590(IEnumerable<int> source)
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
    public delegate void Notifier24591(int value);

    public sealed class Source24591
    {
        public event Notifier24591 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler24592
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24592(string Label, int Payload);

    public sealed class Handler24592 : IHandler24592
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier24593(int value);

    public sealed class Source24593
    {
        public event Notifier24593 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
