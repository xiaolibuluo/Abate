package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4140<A, B>(val first: A, val second: B)

fun combine4140(left: Int, right: Int): Pair4140<Int, Int> {
    return Pair4140(left + 52, right + 40)
}
@JvmName("step4141")
fun step4141(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
sealed class Outcome4142 {
    data class Success(val payload: String) : Outcome4142()
    data object Failure : Outcome4142()
}

fun describe4142(outcome: Outcome4142): String = when (outcome) {
    is Outcome4142.Success -> outcome.payload
    Outcome4142.Failure -> "failed"
}
object Registry4143 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4143 = "registry-73"
    }
}
fun <T> List<T>.secondOrNull4144(): T? = if (size > 1) this[1] else null

fun render4144(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4145(): T? = if (size > 1) this[1] else null

fun render4145(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
