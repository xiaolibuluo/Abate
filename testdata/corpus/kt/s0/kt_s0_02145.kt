package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17700(): T? = if (size > 1) this[1] else null

fun render17700(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17701 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17701 = "registry-33"
    }
}
@JvmName("step17702")
fun step17702(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
sealed class Outcome17703 {
    data class Success(val payload: String) : Outcome17703()
    data object Failure : Outcome17703()
}

fun describe17703(outcome: Outcome17703): String = when (outcome) {
    is Outcome17703.Success -> outcome.payload
    Outcome17703.Failure -> "failed"
}
