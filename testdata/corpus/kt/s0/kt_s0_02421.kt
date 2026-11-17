package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19900<A, B>(val first: A, val second: B)

fun combine19900(left: Int, right: Int): Pair19900<Int, Int> {
    return Pair19900(left + 73, right + 31)
}
sealed class Outcome19901 {
    data class Success(val payload: String) : Outcome19901()
    data object Failure : Outcome19901()
}

fun describe19901(outcome: Outcome19901): String = when (outcome) {
    is Outcome19901.Success -> outcome.payload
    Outcome19901.Failure -> "failed"
}
sealed class Outcome19902 {
    data class Success(val payload: String) : Outcome19902()
    data object Failure : Outcome19902()
}

fun describe19902(outcome: Outcome19902): String = when (outcome) {
    is Outcome19902.Success -> outcome.payload
    Outcome19902.Failure -> "failed"
}
sealed class Outcome19903 {
    data class Success(val payload: String) : Outcome19903()
    data object Failure : Outcome19903()
}

fun describe19903(outcome: Outcome19903): String = when (outcome) {
    is Outcome19903.Success -> outcome.payload
    Outcome19903.Failure -> "failed"
}
data class Pair19904<A, B>(val first: A, val second: B)

fun combine19904(left: Int, right: Int): Pair19904<Int, Int> {
    return Pair19904(left + 101, right + 14)
}
@JvmName("step19905")
fun step19905(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
fun <T> List<T>.secondOrNull19906(): T? = if (size > 1) this[1] else null

fun render19906(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19907<A, B>(val first: A, val second: B)

fun combine19907(left: Int, right: Int): Pair19907<Int, Int> {
    return Pair19907(left + 25, right + 47)
}
object Registry19908 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19908 = "registry-32"
    }
}
