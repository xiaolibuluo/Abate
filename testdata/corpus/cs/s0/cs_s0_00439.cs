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

    public delegate void Notifier4130(int value);

    public sealed class Source4130
    {
        public event Notifier4130 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler4131
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4131(string Label, int Payload);

    public sealed class Handler4131 : IHandler4131
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate4132(IEnumerable<int> source)
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
