package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12030 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12030 = "registry-27"
    }
}
fun <T> List<T>.secondOrNull12031(): T? = if (size > 1) this[1] else null

fun render12031(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12032 {
    data class Success(val payload: String) : Outcome12032()
    data object Failure : Outcome12032()
}

fun describe12032(outcome: Outcome12032): String = when (outcome) {
    is Outcome12032.Success -> outcome.payload
    Outcome12032.Failure -> "failed"
}
@JvmName("step12033")
fun step12033(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome12034 {
    data class Success(val payload: String) : Outcome12034()
    data object Failure : Outcome12034()
}

fun describe12034(outcome: Outcome12034): String = when (outcome) {
    is Outcome12034.Success -> outcome.payload
    Outcome12034.Failure -> "failed"
}
@JvmName("step12035")
fun step12035(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
data class Pair12036<A, B>(val first: A, val second: B)

fun combine12036(left: Int, right: Int): Pair12036<Int, Int> {
    return Pair12036(left + 69, right + 17)
}
