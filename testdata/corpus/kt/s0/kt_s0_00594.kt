package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4820")
fun step4820(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
sealed class Outcome4821 {
    data class Success(val payload: String) : Outcome4821()
    data object Failure : Outcome4821()
}

fun describe4821(outcome: Outcome4821): String = when (outcome) {
    is Outcome4821.Success -> outcome.payload
    Outcome4821.Failure -> "failed"
}
data class Pair4822<A, B>(val first: A, val second: B)

fun combine4822(left: Int, right: Int): Pair4822<Int, Int> {
    return Pair4822(left + 77, right + 50)
}
fun <T> List<T>.secondOrNull4823(): T? = if (size > 1) this[1] else null

fun render4823(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4824<A, B>(val first: A, val second: B)

fun combine4824(left: Int, right: Int): Pair4824<Int, Int> {
    return Pair4824(left + 91, right + 11)
}
object Registry4825 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4825 = "registry-98"
    }
}
sealed class Outcome4826 {
    data class Success(val payload: String) : Outcome4826()
    data object Failure : Outcome4826()
}

fun describe4826(outcome: Outcome4826): String = when (outcome) {
    is Outcome4826.Success -> outcome.payload
    Outcome4826.Failure -> "failed"
}
@JvmName("step4827")
fun step4827(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
fun <T> List<T>.secondOrNull4828(): T? = if (size > 1) this[1] else null

fun render4828(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
