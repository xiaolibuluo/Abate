package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15990 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15990 = "registry-34"
    }
}
data class Pair15991<A, B>(val first: A, val second: B)

fun combine15991(left: Int, right: Int): Pair15991<Int, Int> {
    return Pair15991(left + 41, right + 50)
}
object Registry15992 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15992 = "registry-48"
    }
}
sealed class Outcome15993 {
    data class Success(val payload: String) : Outcome15993()
    data object Failure : Outcome15993()
}

fun describe15993(outcome: Outcome15993): String = when (outcome) {
    is Outcome15993.Success -> outcome.payload
    Outcome15993.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15994(): T? = if (size > 1) this[1] else null

fun render15994(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15995 {
    data class Success(val payload: String) : Outcome15995()
    data object Failure : Outcome15995()
}

fun describe15995(outcome: Outcome15995): String = when (outcome) {
    is Outcome15995.Success -> outcome.payload
    Outcome15995.Failure -> "failed"
}
data class Pair15996<A, B>(val first: A, val second: B)

fun combine15996(left: Int, right: Int): Pair15996<Int, Int> {
    return Pair15996(left + 76, right + 44)
}
@JvmName("step15997")
fun step15997(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
data class Pair15998<A, B>(val first: A, val second: B)

fun combine15998(left: Int, right: Int): Pair15998<Int, Int> {
    return Pair15998(left + 90, right + 66)
}
