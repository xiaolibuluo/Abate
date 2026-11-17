package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19190<A, B>(val first: A, val second: B)

fun combine19190(left: Int, right: Int): Pair19190<Int, Int> {
    return Pair19190(left + 23, right + 22)
}
data class Pair19191<A, B>(val first: A, val second: B)

fun combine19191(left: Int, right: Int): Pair19191<Int, Int> {
    return Pair19191(left + 30, right + 33)
}
sealed class Outcome19192 {
    data class Success(val payload: String) : Outcome19192()
    data object Failure : Outcome19192()
}

fun describe19192(outcome: Outcome19192): String = when (outcome) {
    is Outcome19192.Success -> outcome.payload
    Outcome19192.Failure -> "failed"
}
object Registry19193 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19193 = "registry-44"
    }
}
fun <T> List<T>.secondOrNull19194(): T? = if (size > 1) this[1] else null

fun render19194(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
