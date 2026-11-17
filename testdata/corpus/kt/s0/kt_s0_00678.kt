package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5560(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5561 {
    data class Success(val payload: String) : Outcome5561()
    data object Failure : Outcome5561()
}

fun describe5561(outcome: Outcome5561): String = when (outcome) {
    is Outcome5561.Success -> outcome.payload
    Outcome5561.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5562(): T? = if (size > 1) this[1] else null

fun render5562(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5563 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5563 = "registry-76"
    }
}
data class Pair5564<A, B>(val first: A, val second: B)

fun combine5564(left: Int, right: Int): Pair5564<Int, Int> {
    return Pair5564(left + 83, right + 41)
}
