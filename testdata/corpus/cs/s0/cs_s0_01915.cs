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

    public async Task<int> Aggregate17880(IEnumerable<int> source)
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
    public interface IHandler17881
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17881(string Label, int Payload);

    public sealed class Handler17881 : IHandler17881
    {
        public void Handle(int value) { }
    }
    public interface IHandler17882
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17882(string Label, int Payload);

    public sealed class Handler17882 : IHandler17882
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap17883
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
