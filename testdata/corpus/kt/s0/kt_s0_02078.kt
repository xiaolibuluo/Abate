package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17170(): T? = if (size > 1) this[1] else null

fun render17170(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17171 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17171 = "registry-23"
    }
}
sealed class Outcome17172 {
    data class Success(val payload: String) : Outcome17172()
    data object Failure : Outcome17172()
}

fun describe17172(outcome: Outcome17172): String = when (outcome) {
    is Outcome17172.Success -> outcome.payload
    Outcome17172.Failure -> "failed"
}
