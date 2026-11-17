package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9840(): T? = if (size > 1) this[1] else null

fun render9840(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9841 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9841 = "registry-97"
    }
}
data class Pair9842<A, B>(val first: A, val second: B)

fun combine9842(left: Int, right: Int): Pair9842<Int, Int> {
    return Pair9842(left + 7, right + 26)
}
