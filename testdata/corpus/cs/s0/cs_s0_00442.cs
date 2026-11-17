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

    public interface IHandler4160
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4160(string Label, int Payload);

    public sealed class Handler4160 : IHandler4160
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate4161(IEnumerable<int> source)
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
    public delegate void Notifier4162(int value);

    public sealed class Source4162
    {
        public event Notifier4162 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard4163(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 99;
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
}
