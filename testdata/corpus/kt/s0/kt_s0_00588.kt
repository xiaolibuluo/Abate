package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4770<A, B>(val first: A, val second: B)

fun combine4770(left: Int, right: Int): Pair4770<Int, Int> {
    return Pair4770(left + 95, right + 54)
}
@JvmName("step4771")
fun step4771(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
fun <T> List<T>.secondOrNull4772(): T? = if (size > 1) this[1] else null

fun render4772(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4773 {
    data class Success(val payload: String) : Outcome4773()
    data object Failure : Outcome4773()
}

fun describe4773(outcome: Outcome4773): String = when (outcome) {
    is Outcome4773.Success -> outcome.payload
    Outcome4773.Failure -> "failed"
}
@JvmName("step4774")
fun step4774(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
@JvmName("step4775")
fun step4775(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
fun <T> List<T>.secondOrNull4776(): T? = if (size > 1) this[1] else null

fun render4776(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4777<A, B>(val first: A, val second: B)

fun combine4777(left: Int, right: Int): Pair4777<Int, Int> {
    return Pair4777(left + 47, right + 9)
}
object Registry4778 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4778 = "registry-54"
    }
}
