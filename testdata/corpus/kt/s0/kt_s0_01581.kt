package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13100 {
    data class Success(val payload: String) : Outcome13100()
    data object Failure : Outcome13100()
}

fun describe13100(outcome: Outcome13100): String = when (outcome) {
    is Outcome13100.Success -> outcome.payload
    Outcome13100.Failure -> "failed"
}
sealed class Outcome13101 {
    data class Success(val payload: String) : Outcome13101()
    data object Failure : Outcome13101()
}

fun describe13101(outcome: Outcome13101): String = when (outcome) {
    is Outcome13101.Success -> outcome.payload
    Outcome13101.Failure -> "failed"
}
data class Pair13102<A, B>(val first: A, val second: B)

fun combine13102(left: Int, right: Int): Pair13102<Int, Int> {
    return Pair13102(left + 74, right + 51)
}
@JvmName("step13103")
fun step13103(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
data class Pair13104<A, B>(val first: A, val second: B)

fun combine13104(left: Int, right: Int): Pair13104<Int, Int> {
    return Pair13104(left + 88, right + 12)
}
fun <T> List<T>.secondOrNull13105(): T? = if (size > 1) this[1] else null

fun render13105(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry13106 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13106 = "registry-5"
    }
}
data class Pair13107<A, B>(val first: A, val second: B)

fun combine13107(left: Int, right: Int): Pair13107<Int, Int> {
    return Pair13107(left + 12, right + 45)
}
fun <T> List<T>.secondOrNull13108(): T? = if (size > 1) this[1] else null

fun render13108(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
