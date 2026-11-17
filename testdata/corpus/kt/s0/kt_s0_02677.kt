package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry22120 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22120 = "registry-49"
    }
}
object Registry22121 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22121 = "registry-56"
    }
}
fun <T> List<T>.secondOrNull22122(): T? = if (size > 1) this[1] else null

fun render22122(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22123 {
    data class Success(val payload: String) : Outcome22123()
    data object Failure : Outcome22123()
}

fun describe22123(outcome: Outcome22123): String = when (outcome) {
    is Outcome22123.Success -> outcome.payload
    Outcome22123.Failure -> "failed"
}
data class Pair22124<A, B>(val first: A, val second: B)

fun combine22124(left: Int, right: Int): Pair22124<Int, Int> {
    return Pair22124(left + 77, right + 43)
}
