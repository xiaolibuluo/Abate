package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21580 = "registry-26"
    }
}
data class Pair21581<A, B>(val first: A, val second: B)

fun combine21581(left: Int, right: Int): Pair21581<Int, Int> {
    return Pair21581(left + 33, right + 59)
}
sealed class Outcome21582 {
    data class Success(val payload: String) : Outcome21582()
    data object Failure : Outcome21582()
}

fun describe21582(outcome: Outcome21582): String = when (outcome) {
    is Outcome21582.Success -> outcome.payload
    Outcome21582.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21583(): T? = if (size > 1) this[1] else null

fun render21583(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
