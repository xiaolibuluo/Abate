package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull15010(): T? = if (size > 1) this[1] else null

fun render15010(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step15011")
fun step15011(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
sealed class Outcome15012 {
    data class Success(val payload: String) : Outcome15012()
    data object Failure : Outcome15012()
}

fun describe15012(outcome: Outcome15012): String = when (outcome) {
    is Outcome15012.Success -> outcome.payload
    Outcome15012.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15013(): T? = if (size > 1) this[1] else null

fun render15013(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15014<A, B>(val first: A, val second: B)

fun combine15014(left: Int, right: Int): Pair15014<Int, Int> {
    return Pair15014(left + 49, right + 68)
}
sealed class Outcome15015 {
    data class Success(val payload: String) : Outcome15015()
    data object Failure : Outcome15015()
}

fun describe15015(outcome: Outcome15015): String = when (outcome) {
    is Outcome15015.Success -> outcome.payload
    Outcome15015.Failure -> "failed"
}
data class Pair15016<A, B>(val first: A, val second: B)

fun combine15016(left: Int, right: Int): Pair15016<Int, Int> {
    return Pair15016(left + 63, right + 29)
}
object Registry15017 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15017 = "registry-70"
    }
}
