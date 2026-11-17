package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7160 {
    data class Success(val payload: String) : Outcome7160()
    data object Failure : Outcome7160()
}

fun describe7160(outcome: Outcome7160): String = when (outcome) {
    is Outcome7160.Success -> outcome.payload
    Outcome7160.Failure -> "failed"
}
sealed class Outcome7161 {
    data class Success(val payload: String) : Outcome7161()
    data object Failure : Outcome7161()
}

fun describe7161(outcome: Outcome7161): String = when (outcome) {
    is Outcome7161.Success -> outcome.payload
    Outcome7161.Failure -> "failed"
}
@JvmName("step7162")
fun step7162(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
object Registry7163 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7163 = "registry-22"
    }
}
fun <T> List<T>.secondOrNull7164(): T? = if (size > 1) this[1] else null

fun render7164(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7165<A, B>(val first: A, val second: B)

fun combine7165(left: Int, right: Int): Pair7165<Int, Int> {
    return Pair7165(left + 36, right + 13)
}
@JvmName("step7166")
fun step7166(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
