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

    public interface IHandler19870
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19870(string Label, int Payload);

    public sealed class Handler19870 : IHandler19870
    {
        public void Handle(int value) { }
    }
    public interface IHandler19871
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19871(string Label, int Payload);

    public sealed class Handler19871 : IHandler19871
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier19872(int value);

    public sealed class Source19872
    {
        public event Notifier19872 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale19873<T>(T value) where T : struct
    {
        int number = 55;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public static int Scale19874<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 54; }
        return number;
    }
    public delegate void Notifier19875(int value);

    public sealed class Source19875
    {
        public event Notifier19875 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate19876(IEnumerable<int> source)
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
