package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20330 {
    data class Success(val payload: String) : Outcome20330()
    data object Failure : Outcome20330()
}

fun describe20330(outcome: Outcome20330): String = when (outcome) {
    is Outcome20330.Success -> outcome.payload
    Outcome20330.Failure -> "failed"
}
sealed class Outcome20331 {
    data class Success(val payload: String) : Outcome20331()
    data object Failure : Outcome20331()
}

fun describe20331(outcome: Outcome20331): String = when (outcome) {
    is Outcome20331.Success -> outcome.payload
    Outcome20331.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20332(): T? = if (size > 1) this[1] else null

fun render20332(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20333 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20333 = "registry-71"
    }
}
@JvmName("step20334")
fun step20334(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
sealed class Outcome20335 {
    data class Success(val payload: String) : Outcome20335()
    data object Failure : Outcome20335()
}

fun describe20335(outcome: Outcome20335): String = when (outcome) {
    is Outcome20335.Success -> outcome.payload
    Outcome20335.Failure -> "failed"
}
data class Pair20336<A, B>(val first: A, val second: B)

fun combine20336(left: Int, right: Int): Pair20336<Int, Int> {
    return Pair20336(left + 92, right + 32)
}
