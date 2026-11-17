package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24770<A, B>(val first: A, val second: B)

fun combine24770(left: Int, right: Int): Pair24770<Int, Int> {
    return Pair24770(left + 99, right + 24)
}
@JvmName("step24771")
fun step24771(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
fun <T> List<T>.secondOrNull24772(): T? = if (size > 1) this[1] else null

fun render24772(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24773<A, B>(val first: A, val second: B)

fun combine24773(left: Int, right: Int): Pair24773<Int, Int> {
    return Pair24773(left + 23, right + 57)
}
sealed class Outcome24774 {
    data class Success(val payload: String) : Outcome24774()
    data object Failure : Outcome24774()
}

fun describe24774(outcome: Outcome24774): String = when (outcome) {
    is Outcome24774.Success -> outcome.payload
    Outcome24774.Failure -> "failed"
}
@JvmName("step24775")
fun step24775(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
object Registry24776 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24776 = "registry-44"
    }
}
fun <T> List<T>.secondOrNull24777(): T? = if (size > 1) this[1] else null

fun render24777(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24778 {
    data class Success(val payload: String) : Outcome24778()
    data object Failure : Outcome24778()
}

fun describe24778(outcome: Outcome24778): String = when (outcome) {
    is Outcome24778.Success -> outcome.payload
    Outcome24778.Failure -> "failed"
}
