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

    public interface IHandler19710
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19710(string Label, int Payload);

    public sealed class Handler19710 : IHandler19710
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier19711(int value);

    public sealed class Source19711
    {
        public event Notifier19711 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19712(int value);

    public sealed class Source19712
    {
        public event Notifier19712 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate19713(IEnumerable<int> source)
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
